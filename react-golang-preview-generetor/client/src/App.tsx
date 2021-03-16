import React, { FormEvent, useState, useRef, ChangeEvent } from "react";
import axios from "axios";

interface ImagePreview {
  screenshot: string;
}

export default function App() {
  const [websiteUrl, setWebsiteUrl] = useState("");
  const [thumbnail, setThumbnail] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(false);

  const inputRef = useRef<HTMLInputElement | null>(null);

  async function onSubmit(event: FormEvent) {
    event.preventDefault();

    if (!websiteUrl) {
      inputRef.current?.focus();
      return null;
    }

    setError(false);
    setLoading(true);

    const config = {
      headers: {
        "Content-Type": "application/x-www-form-urlencoded",
      },
    };

    try {
      const { data } = await axios.post<ImagePreview>(
        "http://localhost:8080/api/thumbnail",
        {
          url: websiteUrl,
        },
        config
      );

      setThumbnail(data.screenshot);
      console.log(data);
    } catch (error) {
      setError(true);
    }

    setLoading(false);
  }

  function onChange(event: ChangeEvent<HTMLInputElement>) {
    setWebsiteUrl(event.target.value);
  }

  return (
    <div id="app" className="container">
      <div className="row">
        <div className="content">
          <h1>Generate a thumbnail of a website</h1>

          <form onSubmit={onSubmit}>
            <div className="form-group">
              <input
                type="text"
                aria-label="Enter a website"
                id="website-input"
                placeholder="Enter a website"
                className="form-control"
                ref={inputRef}
                value={websiteUrl}
                onChange={onChange}
              />
            </div>
            <div className="form-group">
              <button type="submit" className="btn">
                Generate!
              </button>
            </div>
          </form>

          {loading && <h3>Loading...</h3>}
          {error && <h3>Error!</h3>}

          {thumbnail && (
            <img src={thumbnail} alt="Preview" className="preview" />
          )}
        </div>
      </div>
      <small className="dev-info">build by Davi Pereira</small>
    </div>
  );
}
