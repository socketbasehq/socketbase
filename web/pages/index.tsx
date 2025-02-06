import { AnalyticsConnections } from '@socketbase/components/analytics/connections';
import { AnalyticsMessages } from '@socketbase/components/analytics/messages';
import LearnMore from '@socketbase/components/learn-more';

function IndexPage() {
  return (
    <div>
      <h1 className="text-xl font-bold">Dashboard</h1>
      <p className="text-sm text-muted-foreground">
        Welcome to your dashboard. Here you can see your app's analytics and
        metrics.
      </p>

      <div className="grid grid-cols-2 gap-4 my-10">
        <AnalyticsConnections />
        <AnalyticsMessages />
      </div>

      <LearnMore />
    </div>
  );
}

export default IndexPage;
