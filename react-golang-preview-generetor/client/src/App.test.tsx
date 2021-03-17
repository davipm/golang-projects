import React, { ReactNode } from "react";
import { render, fireEvent } from "@testing-library/react";
import { QueryClientProvider, QueryClient } from "react-query";
import App from "./App";

const queryClient = new QueryClient();
const Wrapper = ({ children }: { children: ReactNode }) => (
  <QueryClientProvider client={queryClient}>{children}</QueryClientProvider>
);

describe("<App />", () => {
  it("should render the App", () => {
    const { container } = render(
      <Wrapper>
        <App />
      </Wrapper>
    );
    expect(container).toBeTruthy();
  });

  it("should test the input", () => {
    const { getByPlaceholderText, getByText } = render(
      <Wrapper>
        <App />
      </Wrapper>
    );
    const inputMessage: any = getByPlaceholderText("Enter a website");
    const submit: any = getByText("Generate!");

    fireEvent.change(inputMessage, { target: { value: "Text" } });
    expect(inputMessage.value).toBe("Text");

    fireEvent.change(inputMessage, { target: { value: "" } });
    expect(inputMessage.value).toBe("");

    fireEvent.click(submit);
  });
});
