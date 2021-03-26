import styled from "styled-components/macro";

export const Container = styled.div`
  display: block;
  width: 95%;
  margin: auto;

  input {
    display: block;
    width: 100%;
    height: 40px;
    padding: .3rem .7rem;
    font-size: 1.2rem;
    font-weight: 300;
    line-height: 1.5;
    text-align: left;
    color: black;
    background-color: #FFF;
    background-clip: padding-box;
    border: 1px solid #c3c3c3;
    border-radius: .25rem;
    outline: 0;
    transition: all .15s ease-in-out;

    :focus {
      box-shadow: 0 0 6px rgba(0, 0, 0, 0.6);
    }
  }
`;
