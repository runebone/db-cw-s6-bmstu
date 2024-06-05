import pandas as pd

reader = pd.read_csv("cache.csv")
dfcache = pd.DataFrame(reader)

reader = pd.read_csv("nocache.csv")
dfnocache = pd.DataFrame(reader)

gb = ["type", "name"]
# gb = ["type", "name", "rps"]

c = dfcache.groupby(gb).mean().reset_index()
nc = dfnocache.groupby(gb).mean().reset_index()

print(c)
print(nc)
