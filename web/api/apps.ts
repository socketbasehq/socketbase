import { createApi } from '@reduxjs/toolkit/query/react';
import { baseQuery } from './_baseQuery';

type CreateAppRequest = {
  name: string;
};

export type App = {
  id: string;
  name: string;
  app_id: string;
  app_secret: string;
  app_key: string;
  created_at: string;
  updated_at: string;
};

export const appsApi = createApi({
  baseQuery,
  reducerPath: 'appsApi',
  endpoints: builder => ({
    createApp: builder.mutation<App, CreateAppRequest>({
      query: body => ({
        url: '/apps',
        method: 'POST',
        body,
      }),
    }),
    listApps: builder.query<{ data: App[] }, void>({
      query: () => '/apps',
    }),
    getApp: builder.query<{ data: App }, { id: string }>({
      query: ({ id }) => `/apps/${id}`,
    }),
  }),
});

export const { useCreateAppMutation, useListAppsQuery, useGetAppQuery } =
  appsApi;
