import type { HTMLAttributes, ReactNode } from "react";

type DesignCardProps = HTMLAttributes<HTMLElement> & {
  as?: "article" | "div" | "section";
  children: ReactNode;
};

export function DesignCard({
  as: Element = "article",
  className,
  children,
  ...props
}: DesignCardProps) {
  return (
    <Element className={`ed-card${className ? ` ${className}` : ""}`} {...props}>
      {children}
    </Element>
  );
}
