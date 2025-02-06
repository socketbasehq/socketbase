function AppOverview() {
  return (
    <>
      <div className="grid grid-cols-2 gap-4">
        <div className="p-4 rounded-xl border bg-background">
          <h1 className="text-sm text-muted-foreground font-medium">
            Peak connections today
          </h1>
          <h1 className="text-2xl font-semibold mt-4">12</h1>
        </div>
        <div className="p-4 rounded-xl border bg-background">
          <h1 className="text-sm text-muted-foreground font-medium">
            Total messages sent today
          </h1>
          <h1 className="text-2xl font-semibold mt-4">41.4k</h1>
        </div>

        <div></div>
      </div>
    </>
  );
}

export default AppOverview;
