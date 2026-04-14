import fs from 'fs';
import path from 'path';

export function getAwesomeData() {
  // website klasöründen iki üst dizine çıkıp data/data.json'a ulaşıyoruz
  // monorepo mimarisinin en güzel yanlarından biri budur.
  const filePath = path.join(process.cwd(), '../data/data.json');
  
  const fileContents = fs.readFileSync(filePath, 'utf8');
  const data = JSON.parse(fileContents);
  
  return data;
}