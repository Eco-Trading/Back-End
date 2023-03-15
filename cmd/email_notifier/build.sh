cp ../../go.mod .
cp ../../go.sum .
cp -r ../../util ./util
cp -r ../../internal ./internal
cp -r ../../config ./config

docker build \
--compress \
--memory 126m \
--memory-swap 256m \
--build-arg GRPC_PORT=19000 \
--build-arg DATABASE_MONGODB_URL="mongodb://mongo_category_first,mongo_category_secound,mongo_category_third/?replicaSet=replicated_database&maxPoolSize=20&w=majority" \
--tag lukewre/account-manager:v0.1.0 \
.