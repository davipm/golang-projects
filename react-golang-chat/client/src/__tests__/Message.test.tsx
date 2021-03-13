import { render } from "@testing-library/react";
import "@testing-library/jest-dom/extend-expect";

import Message from "../components/message";

describe('<Messgae />', () => {
  it('should render message', () => {
    const message = { type: 1, body: 'Test' }
    const { getByText } = render(<Message message={message} />)

    expect(getByText("Test").textContent).toBe("Test")
  });
});
