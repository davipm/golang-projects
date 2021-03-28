import styled from "styled-components/macro";

export const Container = styled.div`
  display: block;
  width: 95%;
  margin: auto;

  input {
    display: block;
    width: 100%;
    height: 40px;
    padding: 0.3rem 0.7rem;
    font-size: 1.2rem;
    font-weight: 300;
    line-height: 1.5;
    text-align: left;
    color: black;
    background-color: #fff;
    background-clip: padding-box;
    border: 1px solid #c3c3c3;
    border-radius: 0.25rem;
    outline: 0;
    transition: all 0.15s ease-in-out;

    :focus {
      box-shadow: 0 5px 15px -5px rgba(0, 0, 0, 0.2);
    }
  }
`;
