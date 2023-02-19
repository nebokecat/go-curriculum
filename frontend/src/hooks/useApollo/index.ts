import {
  ApolloClient,
  ApolloLink,
  createHttpLink,
  InMemoryCache,
  NormalizedCacheObject,
} from '@apollo/client';
import { useMemo } from 'react';

let apolloClient: ApolloClient<NormalizedCacheObject> | undefined;

const httpLink = createHttpLink({
  uri: 'http://localhost:8080/query',
});

export const useApollo = () => {
  const link = ApolloLink.from([httpLink]);

  return useMemo(() => initializeApollo(link), [link]);
};

const initializeApollo = (link: ApolloLink) => {
  const _apolloClient = apolloClient ?? createApolloClient(link);
  if (!apolloClient) apolloClient = _apolloClient;

  return _apolloClient;
};

const createApolloClient = (link: ApolloLink) => {
  return new ApolloClient({
    link,
    cache: new InMemoryCache(),
  });
};
