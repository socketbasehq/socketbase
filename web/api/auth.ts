import { createApi } from '@reduxjs/toolkit/query/react';
import { baseQuery } from './_baseQuery';

type LoginRequest = {
  username: string;
  password: string;
};

type LoginResponse = {
  token: string;
};

type User = {
  id: string;
  username: string;
};

export const authApi = createApi({
  baseQuery,
  reducerPath: 'authApi',
  tagTypes: ['Auth'],
  endpoints: builder => ({
    getMe: builder.query<User, void>({
      query: () => ({
        url: '/auth/me',
        method: 'GET',
      }),
    }),
    login: builder.mutation<LoginResponse, LoginRequest>({
      query: body => ({
        url: '/auth/login',
        method: 'POST',
        body,
      }),
    }),
  }),
});

export const { useLoginMutation, useGetMeQuery } = authApi;
