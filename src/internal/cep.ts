export function generateBrazilianCEP(): string {
  let digits = randomDigits(8);

  while (digits === '00000000') {
    digits = randomDigits(8);
  }

  return `${digits.slice(0, 5)}-${digits.slice(5)}`;
}

function randomDigits(length: number): string {
  return Array.from({ length }, () => Math.floor(Math.random() * 10)).join('');
}
