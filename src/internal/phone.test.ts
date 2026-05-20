import assert from 'node:assert/strict';
import test from 'node:test';
import { generateBrazilianPhone } from './phone.js';

const PHONE_PATTERN = /^\d{2}9\d{8}$/;
const VALID_DDDS = new Set([
  '11', '12', '13', '14', '15', '16', '17', '18', '19',
  '21', '22', '24', '27', '28',
  '31', '32', '33', '34', '35', '37', '38',
  '41', '42', '43', '44', '45', '46', '47', '48', '49',
  '51', '53', '54', '55',
  '61', '62', '63', '64', '65', '66', '67', '68', '69',
  '71', '73', '74', '75', '77', '79',
  '81', '82', '83', '84', '85', '86', '87', '88', '89',
  '91', '92', '93', '94', '95', '96', '97', '98', '99',
]);

test('generateBrazilianPhone returns DDD-prefixed mobile phone values', () => {
  for (let index = 0; index < 100; index += 1) {
    const phone = generateBrazilianPhone();

    assert.match(phone, PHONE_PATTERN);
    assert.equal(VALID_DDDS.has(phone.slice(0, 2)), true);
  }
});
