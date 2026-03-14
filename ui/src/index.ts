import { LitElement, html, css } from 'lit';
import { customElement, property, state } from 'lit/decorators.js';

/**
 * A demo custom element backed by a Go API.
 * Replace this with your own element implementation.
 *
 * @element core-demo-element
 */
@customElement('core-demo-element')
export class CoreDemoElement extends LitElement {
  /** The API base URL. Defaults to current origin. */
  @property({ type: String, attribute: 'api-url' })
  apiUrl = '';

  @state()
  private message = 'Loading...';

  @state()
  private uptime = '';

  static styles = css`
    :host {
      display: block;
      font-family: system-ui, -apple-system, sans-serif;
      padding: 1.5rem;
      border: 1px solid #e2e8f0;
      border-radius: 0.5rem;
      max-width: 480px;
    }
    h2 {
      margin: 0 0 1rem;
      font-size: 1.25rem;
      color: #1a1b26;
    }
    .status {
      font-size: 0.875rem;
      color: #64748b;
    }
    .badge {
      display: inline-block;
      padding: 0.125rem 0.5rem;
      border-radius: 9999px;
      font-size: 0.75rem;
      font-weight: 600;
      background: #dcfce7;
      color: #166534;
    }
  `;

  connectedCallback() {
    super.connectedCallback();
    this.fetchData();
  }

  private async fetchData() {
    const base = this.apiUrl || window.location.origin;
    try {
      const [helloRes, statusRes] = await Promise.all([
        fetch(`${base}/api/v1/demo/hello`),
        fetch(`${base}/api/v1/demo/status`),
      ]);
      const hello = await helloRes.json();
      const status = await statusRes.json();
      this.message = hello.message;
      this.uptime = status.uptime;
    } catch {
      this.message = 'Failed to connect to API';
    }
  }

  render() {
    return html`
      <h2>${this.message}</h2>
      <div class="status">
        <span class="badge">running</span>
        Uptime: ${this.uptime}
      </div>
    `;
  }
}
