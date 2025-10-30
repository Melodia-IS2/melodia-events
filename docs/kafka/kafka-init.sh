#!/bin/bash
set -e

echo "⏳ Waiting for Kafka to become available..."
while ! (echo > /dev/tcp/kafka/9092) >/dev/null 2>&1; do
  sleep 1
done
echo "✅ Kafka available. Creating topics..."

TOPICS=("artists" "releases" "songs" "events" "__consumer_offsets")

for topic in "${TOPICS[@]}"; do
  echo "Creating topic: $topic"
  /bin/kafka-topics --create --if-not-exists \
    --bootstrap-server kafka:9092 \
    --topic "$topic" \
    --partitions 50 \
    --replication-factor 1
done

echo "📜 Existing topics:"
/bin/kafka-topics --list --bootstrap-server kafka:9092

echo "✅ Initialization complete."
