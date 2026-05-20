import { generateBrazilianCEP } from './internal/cep.js';
import { generateBrazilianCNPJ } from './internal/cnpj.js';
import { generateBrazilianCPF } from './internal/cpf.js';
import { generateBrazilianPhone } from './internal/phone.js';

// TODO(PSF-INTERFACE-1): replace this local contract with
// `import type { ScalarFixtures } from '@lazuli/runtime/scalars'` once the
// sibling runtime interface cell lands.
export interface ScalarFixtureProvider<T = string> {
  generate(): T;
  generateMany?(n: number): T[];
  example?: T;
  invalid?(): T;
}

export type ScalarFixtures = Record<string, ScalarFixtureProvider>;

export const fixtures: ScalarFixtures = {
  BrazilianCPF: {
    generate: () => generateBrazilianCPF(),
    example: '111.444.777-35',
    invalid: () => '000.000.000-00',
  },
  BrazilianCNPJ: {
    generate: () => generateBrazilianCNPJ(),
    example: '11.222.333/0001-81',
    invalid: () => '00.000.000/0000-00',
  },
  BrazilianCEP: {
    generate: () => generateBrazilianCEP(),
    example: '01310-100',
    invalid: () => '0000-0000',
  },
  BrazilianPhone: {
    generate: () => generateBrazilianPhone(),
    example: '11987654321',
    invalid: () => '+55-not-a-phone',
  },
};
