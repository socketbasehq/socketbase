import { clsx, type ClassValue } from 'clsx';
import { twMerge } from 'tailwind-merge';
import generateName from 'boring-name-generator';

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export function generateRandomName() {
  return generateName({ number: true, words: 2 }).dashed;
}
