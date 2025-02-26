import express from 'express';
import { getTableRows, executeAction } from './eosService';

const router = express.Router();

// Health check endpoint
router.get('/health', (req, res) => {
  res.json({ status: 'ok', timestamp: new Date().toISOString() });
});

// Handle contract action requests
router.post('/brdgmng/appaddrmap', async (req, res) => {
  try {
    const { action, data } = req.body;
    
    if (!action || !data) {
      return res.status(400).json({ error: 'Missing required parameters' });
    }
    
    const result = await executeAction(action, data);
    res.json({ success: true, transaction: result });
  } catch (error: any) {
    console.error('Contract execution error:', error);
    res.status(500).json({ error: error.message || 'Contract execution failed' });
  }
});

// Get table data
router.get('/contract/table/:table', async (req, res) => {
  try {
    const { table } = req.params;
    const { scope = '', limit = 10, lower_bound = '', upper_bound = '' } = req.query;
    
    const rows = await getTableRows(
      table, 
      scope as string, 
      Number(limit), 
      lower_bound as string, 
      upper_bound as string
    );
    
    res.json({ rows });
  } catch (error: any) {
    console.error('Error getting table data:', error);
    res.status(500).json({ error: error.message || 'Failed to get table data' });
  }
});

export default router;
