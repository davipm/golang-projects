import { render, fireEvent } from "@testing-library/react";
import "@testing-library/jest-dom/extend-expect";

import ChatInput from "../components/chat-input";

describe('<ChatInput />', () => {
  it('should test input component', () => {
    const onSend = jest.fn()
    const { getByPlaceholderText } = render(<ChatInput onSend={onSend} />)

    const inputMessage: any = getByPlaceholderText("Message")

    fireEvent.change(inputMessage, { target: { value: 'Text' } })
    expect(inputMessage.value).toBe("Text")

    fireEvent.change(inputMessage, { target: { value: '' } })
    expect(inputMessage.value).toBe("")

    fireEvent.keyDown(inputMessage, { key: "Enter", code: "Enter" })
    expect(onSend).toHaveBeenCalledTimes(1)
  });
});
