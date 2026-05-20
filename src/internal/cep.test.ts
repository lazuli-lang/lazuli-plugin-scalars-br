import assert from 'node:assert/strict';
import test from 'node:test';
import { generateBrazilianCEP } from './cep.js';

const CEP_PATTERN = /^\d{5}-\d{3}$/;

test('generateBrazilianCEP returns formatted CEP values', () => {
  for (let index = 0; index < 100; index += 1) {
    const cep = generateBrazilianCEP();

    assert.match(cep, CEP_PATTERN);
    assert.notEqual(cep, '00000-000');
  }
});
