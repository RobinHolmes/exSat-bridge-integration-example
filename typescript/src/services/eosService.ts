import { Api, JsonRpc } from 'eosjs';
import { JsSignatureProvider } from 'eosjs/dist/eosjs-jssig';
import fetch from 'node-fetch';
const { TextDecoder, TextEncoder } = require('text-encoding');
import { config } from '../config';

// Initialize EOS RPC connection
const rpc = new JsonRpc(config.eosEndpoint, { fetch: fetch as any });

// Initialize signature provider if private key is available
const signatureProvider = config.eosPrivateKey ? new JsSignatureProvider([config.eosPrivateKey]) : undefined;

// Initialize EOS API
const api = signatureProvider
  ? new Api({
      rpc,
      signatureProvider,
      textDecoder: new TextDecoder(),
      textEncoder: new TextEncoder(),
    })
  : null;

/**
 * Get table rows data
 */
export async function getTableRows(params: any) {
  try {
    const result = await rpc.get_table_rows(params);
    return result.rows;
  } catch (error) {
    console.error('Error getting table data:', error);
    throw error;
  }
}

/**
 * Execute contract action
 */
export async function executeAction(account: string, name: string, data: any) {
  if (!api) {
    throw new Error('EOS private key not configured, unable to execute transaction');
  }

  const authorization = [
    {
      actor: 'res.xsat',
      permission: 'bridge',
    },
    {
      actor: config.eosAccount,
      permission: 'active',
    },
  ];

  try {
    const result = await api.transact(
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
        blocksBehind: 3,
        expireSeconds: 30,
      }
    );

    return result;
  } catch (error) {
    console.error('Error executing contract action:', error);
    throw error;
  }
}
