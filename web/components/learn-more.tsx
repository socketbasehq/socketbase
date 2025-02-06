import { BookIcon } from 'lucide-react';

const items = [
  {
    title: 'What is Socketbase?',
    link: '/learn-more/what-is-socketbase',
  },
  {
    title: 'How does Socketbase work?',
    link: '/learn-more/how-does-socketbase-work',
  },
  {
    title: 'How do I get started?',
    link: '/learn-more/how-do-i-get-started',
  },
  {
    title: 'How do I get support?',
    link: '/learn-more/how-do-i-get-support',
  },
  {
    title: 'What is the pricing?',
    link: '/learn-more/what-is-the-pricing',
  },
];

function LearnMore() {
  return (
    <div>
      <h1 className="text-xl font-medium">Learn more</h1>
      <div className="grid gap-2 grid-rows-3 grid-cols-2 mt-5">
        {items.map(item => (
          <div key={item.title}>
            <a
              target="_blank"
              href={item.link}
              className="inline-flex gap-2 items-center hover:underline"
            >
              <BookIcon size={20} />
              <h1 className="text-sm">{item.title}</h1>
            </a>
          </div>
        ))}
      </div>
    </div>
  );
}

export default LearnMore;
