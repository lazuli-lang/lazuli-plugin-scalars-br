export function generateBrazilianCPF(): string {
  let baseDigits = randomDigits(9);

  while (allDigitsSame(baseDigits)) {
    baseDigits = randomDigits(9);
  }

  const firstCheckDigit = cpfCheckDigit(baseDigits, 10);
  const secondCheckDigit = cpfCheckDigit([...baseDigits, firstCheckDigit], 11);
  const digits = [...baseDigits, firstCheckDigit, secondCheckDigit].join('');

  return `${digits.slice(0, 3)}.${digits.slice(3, 6)}.${digits.slice(6, 9)}-${digits.slice(9)}`;
}

function cpfCheckDigit(digits: readonly number[], initialWeight: number): number {
  const sum = digits.reduce(
    (total, digit, index) => total + digit * (initialWeight - index),
    0,
  );
  const checkDigit = 11 - (sum % 11);

  return checkDigit >= 10 ? 0 : checkDigit;
}

function randomDigits(length: number): number[] {
  return Array.from({ length }, () => Math.floor(Math.random() * 10));
}

function allDigitsSame(digits: readonly number[]): boolean {
  return digits.every((digit) => digit === digits[0]);
}
