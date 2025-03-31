# NFS Metrics Exporter

Cloud native Prometheus exporter for exporting NFS computation metrics from client side. This tool uses `nfsiostat` to get the status of mounted NFS directories inside a machine. After getting the stats, it will parse them, and export them as Prometheus metrics over `/metrics` endpoint. It also provides health-check metrics that check the status of the exporter itself.
