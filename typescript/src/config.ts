import dotenv from 'dotenv';
import { ContractName } from './utils';

dotenv.config();

export const config = {
  port: process.env.PORT || 3000,
  eosNodeUrl: process.env.EOS_NODE_URL || 'https://rpc-sg.exsat.network',
  eosAccount: process.env.EOS_ACCOUNT || '',
  eosPrivateKey: process.env.EOS_PRIVATE_KEY || '',
  resourcePayment: process.env.RESOURCE_PAYMENT || false,
  btcBridgeContract: ContractName.brdgmng,
  brdgmngPermissionId: process.env.BRDGMNG_PERMISSION_ID || 0,
  multichainBridgeContract: ContractName.cbridge,
};

if (!config.eosAccount || !config.eosPrivateKey) {
  console.error('error: EOS_ACCOUNT and EOS_PRIVATE_KEY environment variables are required');
  process.exit(1);
}
