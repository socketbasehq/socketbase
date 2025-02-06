import { App, useListAppsQuery } from '@socketbase/api/apps';
import { Button } from '@socketbase/components/ui/button';
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from '@socketbase/components/ui/table';
import {
  ColumnDef,
  flexRender,
  getCoreRowModel,
  useReactTable,
} from '@tanstack/react-table';
import { Link } from 'react-router-dom';

const columns: ColumnDef<App>[] = [
  {
    accessorKey: 'id',
    header: 'ID',
  },
  {
    accessorKey: 'name',
    header: 'Name',
  },
  {
    accessorKey: 'app_key',
    header: 'Key',
  },
  {
    accessorKey: 'created_at',
    header: 'Created At',
  },
  {
    id: 'actions',
    cell: ({ row }) => {
      return (
        <div className="flex items-center gap-2">
          <Link to={`/apps/${row.original.id}`}>
            <Button size="sm">View</Button>
          </Link>
        </div>
      );
    },
  },
];

function AppsIndexPage() {
  const { data } = useListAppsQuery();

  const table = useReactTable({
    data: data?.data || [],
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

  return (
    <div>
      <h1 className="text-xl font-bold">Apps</h1>
      <p className="text-sm text-muted-foreground">
        Welcome to your apps. Here you can see your apps and manage them.
      </p>

      <Table className="mt-5 p-8 rounded-xl border bg-background">
        <TableHeader>
          {table.getHeaderGroups().map(headerGroup => (
            <TableRow key={headerGroup.id}>
              {headerGroup.headers.map(header => {
                return (
                  <TableHead key={header.id}>
                    {header.isPlaceholder
                      ? null
                      : flexRender(
                          header.column.columnDef.header,
                          header.getContext(),
                        )}
                  </TableHead>
                );
              })}
            </TableRow>
          ))}
        </TableHeader>
        <TableBody>
          {table.getRowModel().rows?.length ? (
            table.getRowModel().rows.map(row => (
              <TableRow
                key={row.id}
                data-state={row.getIsSelected() && 'selected'}
              >
                {row.getVisibleCells().map(cell => (
                  <TableCell key={cell.id}>
                    {flexRender(cell.column.columnDef.cell, cell.getContext())}
                  </TableCell>
                ))}
              </TableRow>
            ))
          ) : (
            <TableRow>
              <TableCell colSpan={columns.length} className="h-24 text-center">
                No results.
              </TableCell>
            </TableRow>
          )}
        </TableBody>
      </Table>
    </div>
  );
}

export default AppsIndexPage;
