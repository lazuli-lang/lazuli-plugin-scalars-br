import assert from 'node:assert/strict';
import test from 'node:test';
import { generateBrazilianCPF } from './cpf.js';

const CPF_PATTERN = /^\d{3}\.\d{3}\.\d{3}-\d{2}$/;

test('generateBrazilianCPF returns formatted valid CPF values', () => {
  for (let index = 0; index < 100; index += 1) {
    const cpf = generateBrazilianCPF();

    assert.match(cpf, CPF_PATTERN);
    assert.equal(isValidCPF(cpf), true);
  }
});

function isValidCPF(value: string): boolean {
  const digits = value.replace(/\D/g, '').split('').map(Number);

  if (digits.length !== 11 || digits.every((digit) => digit === digits[0])) {
    return false;
  }

  return digits[9] === cpfCheckDigit(digits.slice(0, 9), 10)
    && digits[10] === cpfCheckDigit(digits.slice(0, 10), 11);
}

function cpfCheckDigit(digits: readonly number[], initialWeight: number): number {
  const sum = digits.reduce(
    (total, digit, index) => total + digit * (initialWeight - index),
    0,
  );
  const checkDigit = 11 - (sum % 11);

  return checkDigit >= 10 ? 0 : checkDigit;
}
