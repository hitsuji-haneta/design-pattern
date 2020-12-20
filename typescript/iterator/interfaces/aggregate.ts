import { Iterator } from './iterator.ts';

export interface Aggregate {
  createIterator(): Iterator;
}