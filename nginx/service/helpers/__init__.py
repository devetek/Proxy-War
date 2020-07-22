import time
import prometheus_client
from prometheus_client.core import CollectorRegistry
from prometheus_client import Summary, Counter, Histogram, Gauge

_INF = float("inf")

graphs = {}
graphs['c'] = Counter('python_request_operations_total', 'The total number of processed requests')
graphs['h'] = Histogram('python_request_duration_seconds', 'Histogram for the duration in seconds.', buckets=(1, 2, 5, 6, 10, _INF))

def count_process():
    start = time.time()
    graphs['c'].inc()
    
    time.sleep(0.600)
    end = time.time()
    graphs['h'].observe(end - start)