FROM docker.io/elasticsearch:8.15.0

# Expose Elasticsearch port
EXPOSE 9200

# Set environment variables
ENV discovery.type=single-node
ENV xpack.security.enabled=false
ENV xpack.security.enrollment.enabled=false

# Set JVM heap size (adjust based on your system)
ENV ES_JAVA_OPTS="-Xms512m -Xmx512m"

# Create data directory
RUN mkdir -p /usr/share/elasticsearch/data && \
    chown -R elasticsearch:elasticsearch /usr/share/elasticsearch/data

# Switch to elasticsearch user
USER elasticsearch

# Health check
HEALTHCHECK --interval=30s --timeout=10s --start-period=40s --retries=3 \
  CMD curl -f http://localhost:9200/_cluster/health || exit 1

