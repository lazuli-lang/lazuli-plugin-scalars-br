const FIRST_CHECK_WEIGHTS = [5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2] as const;
const SECOND_CHECK_WEIGHTS = [6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2] as const;

export function generateBrazilianCNPJ(): string {
  const baseDigits = [...randomDigits(8), 0, 0, 0, 1];
  const firstCheckDigit = cnpjCheckDigit(baseDigits, FIRST_CHECK_WEIGHTS);
  const secondCheckDigit = cnpjCheckDigit(
    [...baseDigits, firstCheckDigit],
    SECOND_CHECK_WEIGHTS,
  );
  const digits = [...baseDigits, firstCheckDigit, secondCheckDigit].join('');

  return `${digits.slice(0, 2)}.${digits.slice(2, 5)}.${digits.slice(5, 8)}/${digits.slice(8, 12)}-${digits.slice(12)}`;
}

function cnpjCheckDigit(digits: readonly number[], weights: readonly number[]): number {
  const sum = digits.reduce(
    (total, digit, index) => total + digit * weights[index],
    0,
  );
  const checkDigit = 11 - (sum % 11);

  return checkDigit >= 10 ? 0 : checkDigit;
}

function randomDigits(length: number): number[] {
  return Array.from({ length }, () => Math.floor(Math.random() * 10));
}
