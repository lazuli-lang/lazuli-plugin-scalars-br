import assert from 'node:assert/strict';
import test from 'node:test';
import { generateBrazilianCNPJ } from './cnpj.js';

const CNPJ_PATTERN = /^\d{2}\.\d{3}\.\d{3}\/\d{4}-\d{2}$/;
const FIRST_CHECK_WEIGHTS = [5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2] as const;
const SECOND_CHECK_WEIGHTS = [6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2] as const;

test('generateBrazilianCNPJ returns formatted valid CNPJ values', () => {
  for (let index = 0; index < 100; index += 1) {
    const cnpj = generateBrazilianCNPJ();

    assert.match(cnpj, CNPJ_PATTERN);
    assert.equal(isValidCNPJ(cnpj), true);
  }
});

function isValidCNPJ(value: string): boolean {
  const digits = value.replace(/\D/g, '').split('').map(Number);

  if (digits.length !== 14 || digits.every((digit) => digit === digits[0])) {
    return false;
  }

  return digits[12] === cnpjCheckDigit(digits.slice(0, 12), FIRST_CHECK_WEIGHTS)
    && digits[13] === cnpjCheckDigit(digits.slice(0, 13), SECOND_CHECK_WEIGHTS);
}

function cnpjCheckDigit(digits: readonly number[], weights: readonly number[]): number {
  const sum = digits.reduce(
    (total, digit, index) => total + digit * weights[index],
    0,
  );
  const checkDigit = 11 - (sum % 11);

  return checkDigit >= 10 ? 0 : checkDigit;
}
