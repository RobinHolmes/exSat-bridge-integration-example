import dotenv from 'dotenv';

dotenv.config();

export const config = {
  port: process.env.PORT || 3000,
  eosEndpoint: process.env.EOS_ENDPOINT || 'https://eos.greymass.com',
  eosAccount: process.env.EOS_ACCOUNT || '',
  eosPrivateKey: process.env.EOS_PRIVATE_KEY || '',
  btcBridgeContract: 'brdgmng.xsat',
  multichainBridgeContract: 'cbridge.xsat',
};

if (!config.eosAccount || !config.eosPrivateKey) {
  console.error('error: EOS_ACCOUNT and EOS_PRIVATE_KEY environment variables are required');
  process.exit(1);
}
