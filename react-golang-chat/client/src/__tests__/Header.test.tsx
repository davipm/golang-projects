import { render } from "@testing-library/react";
import "@testing-library/jest-dom/extend-expect";

import Header from "../components/header";

describe('<Header />', () => {
  it('should render component', () => {
    const { container } = render(<Header />)
    expect(container).toBeTruthy()
  });
});
