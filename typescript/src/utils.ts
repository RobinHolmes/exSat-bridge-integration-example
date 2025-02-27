export enum ContractName {
  brdgmng = 'brdgmng.xsat',
  cbridge = 'cbridge.xsat',
  res = 'res.xsat',
}

export enum IndexPosition {
  Primary = 'primary',
  Secondary = 'secondary',
  Tertiary = 'tertiary',
  Fourth = 'fourth',
  Fifth = 'fifth',
  Sixth = 'sixth',
  Seventh = 'seventh',
  Eighth = 'eighth',
  Ninth = 'ninth',
  Tenth = 'tenth',
}

export enum KeyType {
  I64 = 'i64',
  I128 = 'i128',
  I256 = 'i256',
  Float64 = 'float64',
  Float128 = 'float128',
  Ripemd160 = 'ripemd160',
  Sha256 = 'sha256',
  Name = 'name',
}

/**
 * Compute SHA-256 hash
 * @param evmAddress
 */
export function computeId(evmAddress: string): string {
  if (evmAddress.startsWith('0x')) {
    evmAddress = evmAddress.slice(2);
  }
  const result = Buffer.alloc(32);
  result.write(evmAddress, 12, 'hex');
  return result.toString('hex');
}
