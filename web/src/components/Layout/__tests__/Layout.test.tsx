import {render} from '@testing-library/react';
import {MemoryRouter} from 'react-router-dom';
import Layout from '../index';

test('Layout', () => {
  const result = render(
    <MemoryRouter>
      <Layout>
        <h2>This</h2>
      </Layout>
    </MemoryRouter>
  );
  expect(result.container).toMatchSnapshot();
});
