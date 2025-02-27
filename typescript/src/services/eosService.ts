import { Session } from '@wharfkit/session';
import { APIClient } from '@wharfkit/antelope';
import { WalletPluginPrivateKey } from '@wharfkit/wallet-plugin-privatekey';
import { config } from '../config';
import { ContractName } from '../utils';

const rpc = new APIClient({ url: config.eosNodeUrl });
let session: Session | null = null;

async function initializeSession() {
  if (!session) {
    const info = await rpc.v1.chain.get_info();
    session = new Session({
      chain: {
        id: info.chain_id,
        url: config.eosNodeUrl,
      },
      actor: config.eosAccount,
      permission: 'active',
      walletPlugin: new WalletPluginPrivateKey(config.eosPrivateKey),
    });
    // Log session initialization details
    console.log('Session initialized with:', {
      eosNodeUrl: config.eosNodeUrl,
      eosAccount: config.eosAccount,
    });
  }
}

/**
 * Execute action
 */
export async function executeAction(account: string, name: string, data: any) {
  await initializeSession();

  if (!session) {
    throw new Error('Session not initialized');
  }
  const authorization = [
    {
      actor: config.eosAccount,
      permission: 'active',
    },
  ];
  if (config.resourcePayment) {
    authorization.unshift({
      actor: ContractName.res,
      permission: 'bridge',
    });
  }

  try {
    const result = await session.transact(
      {
        actions: [
          {
            account,
            name,
            authorization,
            data,
          },
        ],
      },
      {
        expireSeconds: 30,
      }
    );
    console.log('Transaction was successful boardcasted, transactionId:', result.response?.transaction_id);
    return result;
  } catch (error) {
    console.error('Error executing contract action:', error);
    throw error;
  }
}

/**
 * Get table rows data
 */
export async function getTableRows(params: any) {
  try {
    const result = await rpc.v1.chain.get_table_rows(params);
    return result.rows;
  } catch (error) {
    console.error('Error getting table data:', error);
    throw error;
  }
}
