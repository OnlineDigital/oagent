import { SolidMarkdown } from "solid-markdown";
import remarkGfm from "remark-gfm";

export default function Markdown(props: { text?: string; class?: string }) {
  return (
    <div classList={{ markdown: true, [props.class ?? ""]: !!props.class }}>
      <SolidMarkdown
        skipHtml
        renderingStrategy="memo"
        remarkPlugins={[remarkGfm]}
        children={props.text ?? ""}
      />
    </div>
  );
}
