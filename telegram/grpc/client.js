// Get gRPC packages
const grpc = require("@grpc/grpc-js");
const protoLoader = require("@grpc/proto-loader");

// ProtoLoader
const packageDefinition = protoLoader.loadSync(process.env.TELEGRAM_PROTO_PATH, {
  keepCase: true,
  longs: String,
  enums: String,
  defaults: true,
  oneofs: true,
});

// Load service
const BotService = grpc.loadPackageDefinition(packageDefinition).bot.BotService;

// Create new client
const client = new BotService(
  process.env.TELEGRAM_GRPC_PATH,
  grpc.credentials.createInsecure()
);

// Export client
module.exports = client;