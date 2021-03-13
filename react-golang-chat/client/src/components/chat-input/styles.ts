import styled from "styled-components/macro";

export const Container = styled.div`
  width: 95%;
  display: block;
  margin: auto;

  input {
    padding: 10px;
    margin: 0;
    font-size: 16px;
    border-radius: 5px;
    border: 1px solid rgba(0, 0, 0, 0.1);
    width: 98%;
    box-shadow: 0 5px 15px -5px rgba(0, 0, 0, 0.1);
    outline: 0;
    transition: all 0.15s ease-in-out;

    :focus {
      box-shadow: 0 0 6px rgba(0, 0, 0, 0.6);
    }
  }
`;
