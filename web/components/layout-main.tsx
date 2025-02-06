import { Link, Outlet } from 'react-router-dom';
import { Button } from '@socketbase/components/ui/button';

import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from '@socketbase/components/ui/dialog';
import { Input } from '@socketbase/components/ui/input';
import { Form, FormField } from '@socketbase/components/ui/form';
import { generateRandomName } from '@socketbase/lib/utils';
import { z } from 'zod';
import { useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { useCreateAppMutation } from '@socketbase/api/apps';
import { toast } from 'sonner';

const appCreateSchema = z.object({
  name: z.string().min(1),
});

function MainLayout() {
  const [createApp, { isLoading }] = useCreateAppMutation();
  const form = useForm<z.infer<typeof appCreateSchema>>({
    resolver: zodResolver(appCreateSchema),
    defaultValues: {
      name: generateRandomName(),
    },
  });

  function handleSubmit(values: z.infer<typeof appCreateSchema>) {
    createApp(values)
      .unwrap()
      .then(() => {
        toast.success('App created successfully');
        window.location.reload();
      });
  }

  return (
    <div className="flex h-screen w-screen flex-col bg-neutral-50">
      <div className="h-20 flex items-center border-b bg-background px-10">
        <div className="max-w-5xl flex items-center justify-between gap-4 mx-auto w-full">
          <Link to="/">
            <img
              src="/logo.svg"
              alt="socketbase"
              className="h-12 w-12 rounded-xl"
              draggable={false}
            />
          </Link>
          <div className="flex items-center gap-2">
            <Link to="/apps">
              <Button variant="ghost">Apps</Button>
            </Link>
            <Link to="/stats">
              <Button variant="ghost">Stats</Button>
            </Link>
            <Link to="/docs">
              <Button variant="ghost">Docs</Button>
            </Link>
            <Form {...form}>
              <Dialog>
                <DialogTrigger asChild>
                  <Button>Create App</Button>
                </DialogTrigger>
                <DialogContent>
                  <form
                    onSubmit={form.handleSubmit(handleSubmit)}
                    className="flex flex-col gap-4"
                  >
                    <DialogHeader>
                      <DialogTitle>Create App</DialogTitle>
                      <DialogDescription>
                        Enter the name of your app and click create to get
                        started.
                      </DialogDescription>
                    </DialogHeader>
                    <FormField
                      control={form.control}
                      name="name"
                      render={({ field }) => (
                        <Input placeholder="App Name" {...field} />
                      )}
                    />
                    <DialogFooter>
                      <DialogClose asChild>
                        <Button type="button" variant="outline">
                          Cancel
                        </Button>
                      </DialogClose>
                      <Button type="submit" disabled={isLoading}>
                        Create
                      </Button>
                    </DialogFooter>
                  </form>
                </DialogContent>
              </Dialog>
            </Form>
          </div>
        </div>
      </div>
      <div className="flex-1 max-w-5xl mx-auto w-full py-6">
        <Outlet />
      </div>
    </div>
  );
}

export default MainLayout;
