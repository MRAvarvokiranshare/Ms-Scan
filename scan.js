import fetch from "node-fetch";
import fs from "fs";

const sites = JSON.parse(fs.readFileSync("sites.json"));
const username = process.argv[2];

if (!username) {
  console.log("Usage: node scan.js <username>");
  process.exit(1);
}

for (const site of sites) {
  const url = site.url.replace("{u}", username);
  try {
    const res = await fetch(url, { method: "HEAD" });
    console.log(`${site.name}: ${res.status === 200 ? "✅" : "❌"} ${url}`);
  } catch {
    console.log(`${site.name}: ❌ ${url}`);
  }
}
