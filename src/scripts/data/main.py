import matplotlib.pyplot as plt
import pandas as pd

cache_df = pd.read_csv('cache.csv')
nocache_df = pd.read_csv('nocache.csv')

plt.plot(cache_df['rps'], cache_df['avg_ms'], label='С кешированием', marker='o')
plt.plot(nocache_df['rps'], nocache_df['avg_ms'], label='Без кеширования', marker='x')
plt.xlabel('Количество запросов в секунду')
plt.ylabel('Среднее время ответа (мс)')
plt.legend()
plt.grid(True)
plt.yscale('log')
plt.show()

plt.plot(cache_df['rps'], cache_df['95%'], label='С кешированием', marker='o')
plt.plot(nocache_df['rps'], nocache_df['95%'], label='Без кеширования', marker='x')
plt.xlabel('Количество запросов в секунду')
plt.ylabel('95-й процентиль времени ответа (мс)')
plt.legend()
plt.grid(True)
plt.yscale('log')
plt.show()
