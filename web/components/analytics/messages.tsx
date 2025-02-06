import { CartesianGrid, Line, LineChart, XAxis } from 'recharts';

import {
  ChartConfig,
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
} from '@socketbase/components/ui/chart';

const chartData = [
  { hour: '00:00', count: 100 },
  { hour: '01:00', count: 200 },
  { hour: '02:00', count: 100 },
  { hour: '03:00', count: 90 },
  { hour: '04:00', count: 40 },
  { hour: '05:00', count: 120 },
];

const chartConfig = {
  count: {
    label: 'Messages',
    color: 'hsl(var(--chart-1))',
  },
} satisfies ChartConfig;

export function AnalyticsMessages() {
  return (
    <div>
      <h1 className="text-sm font-medium">Messages sent today</h1>
      <div className="bg-background border rounded-xl p-8 mt-4">
        <ChartContainer config={chartConfig}>
          <LineChart accessibilityLayer data={chartData}>
            <CartesianGrid vertical={false} />
            <XAxis dataKey="hour" tickMargin={8} />
            <ChartTooltip
              cursor={false}
              content={<ChartTooltipContent indicator="dot" />}
            />
            <Line
              dataKey="count"
              type="natural"
              fill="var(--color-count)"
              fillOpacity={0.4}
              stroke="var(--color-count)"
            />
          </LineChart>
        </ChartContainer>
      </div>
    </div>
  );
}
