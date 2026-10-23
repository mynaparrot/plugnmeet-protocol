import { defineConfig } from 'vite';

export default defineConfig({
  build: {
    emptyOutDir: true,
    sourcemap: true,
    lib: {
      entry: 'src/index.ts',
      formats: ['es', 'cjs'],
      fileName: (format) => (format === 'es' ? 'index.mjs.js' : 'index.cjs.js'),
    },
    rolldownOptions: {
      external: [/@bufbuild\/protobuf/],
    },
  },
});
