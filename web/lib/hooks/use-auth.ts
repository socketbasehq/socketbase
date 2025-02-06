import { useGetMeQuery } from '@socketbase/api/auth';

export function useAuth() {
  const { data, isLoading, error } = useGetMeQuery();

  return {
    id: data?.id,
    username: data?.username,
    isLoading,
    error,
  };
}
