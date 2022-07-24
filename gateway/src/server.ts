import "reflect-metadata";
import Express from "express";
import cors from "cors";
import { ApolloServer } from "apollo-server-express";
import { buildSchema } from "type-graphql";
import { HelloResolver } from "./resolvers/HelloResolver";

const main = async () => {
  const app = Express();

  //   MIDDLEWARES
  app.use(cors({ credentials: true }));

  //   BUIDLING TYPEGRAPHQL SCHEMA
  const schema = await buildSchema({
    resolvers: [HelloResolver],
    validate: false,
  });

  // APOLLO STUFFS
  const apollo = new ApolloServer({
    schema,
  });
  await apollo.start();
  apollo.applyMiddleware({ app });

  // START THE TING
  app.listen({ port: 5050 }, () => {
    console.log(
      "GRAPHQL SERVER IS ACTIVE ON http:localhost:5050" + apollo.graphqlPath
    );
  });
};

main().catch(e => console.log(e.message));
