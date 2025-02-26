import express from 'express';
import { config } from './config';
import apiRoutes from './services/apiService';

const app = express();
app.use(express.json());

app.use((req, res, next) => {
  res.header('Access-Control-Allow-Origin', '*');
  res.header('Access-Control-Allow-Headers', 'Origin, X-Requested-With, Content-Type, Accept');
  res.header('Access-Control-Allow-Methods', 'GET, POST, PUT, DELETE, OPTIONS');
  
  if (req.method === 'OPTIONS') {
    res.sendStatus(200);
  } else {
    next();
  }
});

app.use('/api', apiRoutes);

app.get('/', (req, res) => {
  res.json({
    name: 'ExSat Bridge Integration API',
    version: '1.0.0',
    endpoints: {
      contractAction: '/api/brdgmng/appaddrmap',
      getTable: '/api/contract/table/:table'
    }
  });
});

app.listen(config.port, () => {
  console.log(`Server started: http://localhost:${config.port}`);
  console.log(`EOS Account: ${config.eosAccount}`);
});