import { configureStore } from '@reduxjs/toolkit';
import { authApi } from '@socketbase/api/auth';
import { appsApi } from '@socketbase/api/apps';
export const store = configureStore({
  reducer: {
    [authApi.reducerPath]: authApi.reducer,
    [appsApi.reducerPath]: appsApi.reducer,
  },
  middleware: getDefaultMiddleware =>
    getDefaultMiddleware().concat(authApi.middleware, appsApi.middleware),
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
