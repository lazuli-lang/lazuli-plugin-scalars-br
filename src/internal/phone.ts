const COMMON_DDDS = ['11', '21', '31', '41', '51', '61', '71', '81', '85', '91'] as const;

export function generateBrazilianPhone(): string {
  const ddd = COMMON_DDDS[Math.floor(Math.random() * COMMON_DDDS.length)];
  const localNumber = `9${randomDigits(8)}`;

  return `${ddd}${localNumber}`;
}

function randomDigits(length: number): string {
  return Array.from({ length }, () => Math.floor(Math.random() * 10)).join('');
}
