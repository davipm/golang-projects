import React from "react";
import { render, fireEvent } from "@testing-library/react";
import App from "./App";

describe("<App />", () => {
  it("should render the App", () => {
    const { container } = render(<App />);
    expect(container).toBeTruthy();
  });

  it("should test the input", () => {
    const { getByPlaceholderText, getByText } = render(<App />);
    const inputMessage: any = getByPlaceholderText("Enter a website");
    const submit: any = getByText("Generate!");

    fireEvent.change(inputMessage, { target: { value: "Text" } });
    expect(inputMessage.value).toBe("Text");

    fireEvent.change(inputMessage, { target: { value: "" } });
    expect(inputMessage.value).toBe("");

    fireEvent.click(submit);
  });
});
