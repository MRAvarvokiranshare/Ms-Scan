import asyncio, aiohttp, json
from rich.console import Console

async def check_site(session, site, username):
    url = site["url"].format(u=username)
    try:
        async with session.head(url, timeout=5) as r:
            return site["name"], url, r.status == 200
    except:
        return site["name"], url, False

async def main(username):
    sites = json.load(open("sites.json"))
    console = Console()
    async with aiohttp.ClientSession() as session:
        tasks = [check_site(session, s, username) for s in sites]
        results = await asyncio.gather(*tasks)
    for name, url, exists in results:
        icon = "✅" if exists else "❌"
        console.print(f"[bold]{name}[/bold]: {icon} {url}")

if __name__ == "__main__":
    import sys
    if len(sys.argv) != 2:
        print("Usage: python scan.py <username>")
    else:
        asyncio.run(main(sys.argv[1]))
