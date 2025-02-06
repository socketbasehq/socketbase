import { Button } from '@socketbase/components/ui/button';
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@socketbase/components/ui/card';
import { Input } from '@socketbase/components/ui/input';
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from '@socketbase/components/ui/form';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { zodResolver } from '@hookform/resolvers/zod';
import { useLoginMutation } from '@socketbase/api/auth';
import { toast } from 'sonner';
import { useNavigate } from 'react-router-dom';

const loginSchema = z.object({
  username: z.string().min(1),
  password: z.string().min(1),
});

type LoginForm = z.infer<typeof loginSchema>;

function LoginPage() {
  const navigate = useNavigate();
  const [login, { isLoading }] = useLoginMutation();

  const form = useForm<LoginForm>({
    resolver: zodResolver(loginSchema),
    defaultValues: {
      username: '',
      password: '',
    },
  });

  const handleSubmit = (data: LoginForm) => {
    login(data)
      .unwrap()
      .then(() => {
        toast.success('Login successful');
        navigate('/');
      })
      .catch(err => {
        toast.error(err.data.error ?? 'An error occurred');
      });
  };

  return (
    <div className="flex flex-col items-center justify-center h-screen bg-neutral-50">
      <Card className="w-96">
        <CardHeader>
          <img
            src="/logo.svg"
            alt="Socketbase"
            className="w-14 h-14 rounded-2xl mx-auto"
          />
          <CardTitle className="text-center mt-4">Welcome back!</CardTitle>
          <CardDescription className="text-center">
            Enter your username and password to login
          </CardDescription>
        </CardHeader>
        <CardContent>
          <Form {...form}>
            <form
              className="flex flex-col gap-4"
              onSubmit={form.handleSubmit(handleSubmit)}
            >
              <FormField
                control={form.control}
                name="username"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Username</FormLabel>
                    <FormControl className="mt-1">
                      <Input id="username" type="text" {...field} />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="password"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Password</FormLabel>
                    <FormControl className="mt-1">
                      <Input id="password" type="password" {...field} />
                    </FormControl>
                  </FormItem>
                )}
              />

              <Button className="w-full" type="submit" disabled={isLoading}>
                Login
              </Button>
            </form>
          </Form>
        </CardContent>
      </Card>
    </div>
  );
}

export default LoginPage;
