import { render, fireEvent } from "@testing-library/react";
import "@testing-library/jest-dom/extend-expect";

import ChatInput from "../components/chat-input";

describe('<ChatInput />', () => {
  it('should test inout component', () => {
    const onSend = jest.fn()
    const { getByPlaceholderText } = render(<ChatInput send={onSend} />)

    const inputMessage: any = getByPlaceholderText("Message")

    fireEvent.change(inputMessage, { target: { value: 'Text' } })
    expect(inputMessage.value).toBe("Text")

    fireEvent.change(inputMessage, { target: { value: '' } })
    expect(inputMessage.value).toBe("")
  });
});
