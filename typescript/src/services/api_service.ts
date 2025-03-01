import express from 'express';
import { getTableRows, executeAction } from './eos_service';
import { config } from '../config';
import { computeId, IndexPosition, KeyType } from '../utils';

const router = express.Router();

// Health check endpoint
router.get('/health', (req, res) => {
  res.json({ status: 'ok', timestamp: new Date().toISOString() });
});

/**
 * Apply for a corresponding BTC deposit address for an EVM address
 * The EVM address + remark must be unique
 */
router.post('/brdgmng/appaddrmap', async (req, res) => {
  try {
    const { recipient_address, remark } = req.body;

    if (!recipient_address) {
      return res.status(400).json({ error: 'Missing required parameters: recipient_address' });
    }
    const data = {
      actor: config.eosAccount,
      permission_id: config.brdgmngPermissionId,
      recipient_address,
      remark: remark || '',
      assign_deposit_address: null,
    };
    console.log('data', data);
    const result = await executeAction('brdgmng.xsat', 'appaddrmap', data);
    res.json({ success: true, transaction: result });
  } catch (error: any) {
    console.error('Contract execution error:', error);
    res.status(500).json({ error: error.message || 'Contract execution failed' });
  }
});

// Get btc deposit address corresponding to the evm address
router.get('/brdgmng/deposit-address/:recipientAddress', async (req, res) => {
  try {
    const { recipientAddress } = req.params;
    const key = computeId(recipientAddress);
    const params = {
      json: true,
      code: config.btcBridgeContract,
      scope: config.brdgmngPermissionId,
      table: 'addrmappings',
      index_position: IndexPosition.Tertiary,
      key_type: KeyType.Sha256,
      lower_bound: key,
      upper_bound: key,
      limit: 1,
    };
    const rows = await getTableRows(params);
    if (rows && rows.length > 0) {
      return res.json({ recipientAddress, depositAddress: rows[0].btc_address });
    } else {
      return res.status(404).json({ error: 'Deposit address not found' });
    }
  } catch (error: any) {
    console.error('Error getting table data:', error);
    res.status(500).json({ error: error.message || 'Failed to get table data' });
  }
});

export default router;
