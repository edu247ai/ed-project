"use client";

import { useCallback, useEffect, useState } from "react";
import Link from "next/link";
import Image from "next/image";
import {
  Activity,
  AlertCircle,
  ArrowDownRight,
  ArrowRight,
  ArrowUpRight,
  BookOpen,
  Boxes,
  BriefcaseBusiness,
  ChartNoAxesCombined,
  ChevronDown,
  Code2,
  FileCheck2,
  GraduationCap,
  LayoutDashboard,
  LoaderCircle,
  Menu,
  Package,
  RefreshCw,
  Search,
  Settings,
  Shield,
  ShoppingBag,
  Users,
  X,
  type LucideIcon,
} from "lucide-react";
import { DesignCard } from "../../../components/DesignCard";

export type DashboardRole =
  | "student"
  | "teacher"
  | "vendor"
  | "developer"
  | "manager"
  | "superadmin";

const basePath = process.env.NEXT_PUBLIC_BASE_PATH ?? "";

type DashboardSize = "full" | "half" | "third" | "quarter";

type DashboardMetric = {
  id: string;
  label: string;
  value: string | number;
  change?: string;
  trend?: "up" | "down" | "neutral";
};

type DashboardItem = {
  title: string;
  subtitle?: string;
  status?: string;
  value?: string | number;
};

type DashboardProgressItem = {
  label: string;
  value: number;
  max?: number;
  valueLabel?: string;
};

type DashboardSummaryItem = {
  label: string;
  value: string | number;
  description?: string;
};

type DashboardSeries = {
  name: string;
  color?: string;
  values: { label: string; value: number }[];
};

type DashboardWidget =
  | {
      id: string;
      title: string;
      kind: "list";
      size?: DashboardSize;
      items: DashboardItem[];
    }
  | {
      id: string;
      title: string;
      kind: "table";
      size?: DashboardSize;
      columns: string[];
      rows: Record<string, string | number | null>[];
    }
  | {
      id: string;
      title: string;
      kind: "progress";
      size?: DashboardSize;
      items: DashboardProgressItem[];
    }
  | {
      id: string;
      title: string;
      kind: "summary";
      size?: DashboardSize;
      items: DashboardSummaryItem[];
    }
  | {
      id: string;
      title: string;
      kind: "chart";
      size?: DashboardSize;
      series: DashboardSeries[];
    };

type DashboardPayload = {
  role: DashboardRole;
  user?: { name?: string; roleLabel?: string };
  metrics: DashboardMetric[];
  sections: DashboardWidget[];
};

type NavItem = { label: string; icon: LucideIcon };

const roleDetails: Record<
  DashboardRole,
  { title: string; roleLabel: string; nav: NavItem[] }
> = {
  student: {
    title: "Student Dashboard",
    roleLabel: "Student",
    nav: [
      { label: "Dashboard", icon: LayoutDashboard },
      { label: "My Courses", icon: BookOpen },
      { label: "Assignments", icon: FileCheck2 },
      { label: "AI Tutor", icon: Activity },
      { label: "Resources", icon: Boxes },
      { label: "Achievements", icon: GraduationCap },
      { label: "Settings", icon: Settings },
    ],
  },
  teacher: {
    title: "Teacher Dashboard",
    roleLabel: "Teacher",
    nav: [
      { label: "Dashboard", icon: LayoutDashboard },
      { label: "My Courses", icon: BookOpen },
      { label: "Classes", icon: Users },
      { label: "Assignments & Quizzes", icon: FileCheck2 },
      { label: "Submissions", icon: Activity },
      { label: "Content Library", icon: Boxes },
      { label: "Analytics", icon: ChartNoAxesCombined },
      { label: "Settings", icon: Settings },
    ],
  },
  vendor: {
    title: "Vendor Dashboard",
    roleLabel: "Vendor",
    nav: [
      { label: "Dashboard", icon: LayoutDashboard },
      { label: "Products", icon: Package },
      { label: "Inventory", icon: Boxes },
      { label: "Orders", icon: ShoppingBag },
      { label: "Delivery", icon: Activity },
      { label: "Payments", icon: BriefcaseBusiness },
      { label: "Analytics", icon: ChartNoAxesCombined },
      { label: "Settings", icon: Settings },
    ],
  },
  developer: {
    title: "Developer Dashboard",
    roleLabel: "Developer",
    nav: [
      { label: "Overview", icon: LayoutDashboard },
      { label: "Applications", icon: Boxes },
      { label: "API Keys", icon: Shield },
      { label: "Documentation", icon: BookOpen },
      { label: "Issues", icon: FileCheck2 },
      { label: "Contributions", icon: Code2 },
      { label: "Service Health", icon: Activity },
      { label: "Settings", icon: Settings },
    ],
  },
  manager: {
    title: "Manager Dashboard",
    roleLabel: "Manager",
    nav: [
      { label: "Dashboard", icon: LayoutDashboard },
      { label: "People", icon: Users },
      { label: "Courses", icon: BookOpen },
      { label: "Engagement", icon: Activity },
      { label: "Content Review", icon: FileCheck2 },
      { label: "Reports", icon: ChartNoAxesCombined },
      { label: "Support", icon: BriefcaseBusiness },
      { label: "Settings", icon: Settings },
    ],
  },
  superadmin: {
    title: "Super Admin Dashboard",
    roleLabel: "Super Admin",
    nav: [
      { label: "Overview", icon: LayoutDashboard },
      { label: "Users & Roles", icon: Users },
      { label: "Institutions", icon: BriefcaseBusiness },
      { label: "Analytics", icon: ChartNoAxesCombined },
      { label: "Moderation", icon: FileCheck2 },
      { label: "Security & Audit", icon: Shield },
      { label: "System Health", icon: Activity },
      { label: "Settings", icon: Settings },
    ],
  },
};

function isRecord(value: unknown): value is Record<string, unknown> {
  return typeof value === "object" && value !== null && !Array.isArray(value);
}

function isStringOrNumber(value: unknown): value is string | number {
  return typeof value === "string" || typeof value === "number";
}

function parseDashboardPayload(value: unknown, role: DashboardRole): DashboardPayload {
  if (!isRecord(value) || value.role !== role || !Array.isArray(value.metrics) || !Array.isArray(value.sections)) {
    throw new Error("The dashboard API returned data in an unexpected format.");
  }

  const metrics = value.metrics.map((metric): DashboardMetric => {
    if (!isRecord(metric) || typeof metric.id !== "string" || typeof metric.label !== "string" || !isStringOrNumber(metric.value)) {
      throw new Error("The dashboard API returned an invalid metric.");
    }

    if (metric.trend !== undefined && !["up", "down", "neutral"].includes(String(metric.trend))) {
      throw new Error("The dashboard API returned an invalid metric trend.");
    }

    return {
      id: metric.id,
      label: metric.label,
      value: metric.value,
      ...(typeof metric.change === "string" ? { change: metric.change } : {}),
      ...(metric.trend === "up" || metric.trend === "down" || metric.trend === "neutral" ? { trend: metric.trend } : {}),
    };
  });

  const sections = value.sections.map((section): DashboardWidget => {
    if (!isRecord(section) || typeof section.id !== "string" || typeof section.title !== "string") {
      throw new Error("The dashboard API returned an invalid section.");
    }

    const size = section.size;
    if (size !== undefined && !["full", "half", "third", "quarter"].includes(String(size))) {
      throw new Error(`The dashboard API returned an unsupported layout size for "${section.title}".`);
    }

    const dashboardSize: DashboardSize | undefined =
      size === "full" || size === "half" || size === "third" || size === "quarter"
        ? size
        : undefined;
    const common: { id: string; title: string; size?: DashboardSize } = {
      id: section.id,
      title: section.title,
      ...(dashboardSize ? { size: dashboardSize } : {}),
    };

    if (section.kind === "list" && Array.isArray(section.items)) {
      const items = section.items.map((item): DashboardItem => {
        if (!isRecord(item) || typeof item.title !== "string") {
          throw new Error(`The dashboard API returned an invalid item in "${section.title}".`);
        }
        return {
          title: item.title,
          ...(typeof item.subtitle === "string" ? { subtitle: item.subtitle } : {}),
          ...(typeof item.status === "string" ? { status: item.status } : {}),
          ...(isStringOrNumber(item.value) ? { value: item.value } : {}),
        };
      });
      return { ...common, kind: "list", items };
    }

    if (section.kind === "table" && Array.isArray(section.columns) && Array.isArray(section.rows)) {
      if (!section.columns.every((column) => typeof column === "string")) {
        throw new Error(`The dashboard API returned invalid columns for "${section.title}".`);
      }
      const rows = section.rows.map((row) => {
        if (!isRecord(row)) throw new Error(`The dashboard API returned an invalid table row in "${section.title}".`);
        return Object.fromEntries(
          Object.entries(row).map(([key, cell]) => {
            if (cell !== null && !isStringOrNumber(cell)) {
              throw new Error(`The dashboard API returned an invalid cell in "${section.title}".`);
            }
            return [key, cell];
          }),
        ) as Record<string, string | number | null>;
      });
      return { ...common, kind: "table", columns: section.columns as string[], rows };
    }

    if (section.kind === "progress" && Array.isArray(section.items)) {
      const items = section.items.map((item): DashboardProgressItem => {
        if (!isRecord(item) || typeof item.label !== "string" || typeof item.value !== "number") {
          throw new Error(`The dashboard API returned invalid progress data in "${section.title}".`);
        }
        if (item.max !== undefined && typeof item.max !== "number") {
          throw new Error(`The dashboard API returned an invalid progress maximum in "${section.title}".`);
        }
        return {
          label: item.label,
          value: item.value,
          ...(typeof item.max === "number" ? { max: item.max } : {}),
          ...(typeof item.valueLabel === "string" ? { valueLabel: item.valueLabel } : {}),
        };
      });
      return { ...common, kind: "progress", items };
    }

    if (section.kind === "summary" && Array.isArray(section.items)) {
      const items = section.items.map((item): DashboardSummaryItem => {
        if (!isRecord(item) || typeof item.label !== "string" || !isStringOrNumber(item.value)) {
          throw new Error(`The dashboard API returned an invalid summary in "${section.title}".`);
        }
        return {
          label: item.label,
          value: item.value,
          ...(typeof item.description === "string" ? { description: item.description } : {}),
        };
      });
      return { ...common, kind: "summary", items };
    }

    if (section.kind === "chart" && Array.isArray(section.series)) {
      const series = section.series.map((entry): DashboardSeries => {
        if (!isRecord(entry) || typeof entry.name !== "string" || !Array.isArray(entry.values)) {
          throw new Error(`The dashboard API returned an invalid chart series in "${section.title}".`);
        }
        const values = entry.values.map((point) => {
          if (!isRecord(point) || typeof point.label !== "string" || typeof point.value !== "number") {
            throw new Error(`The dashboard API returned an invalid chart point in "${section.title}".`);
          }
          return { label: point.label, value: point.value };
        });
        return {
          name: entry.name,
          values,
          ...(typeof entry.color === "string" ? { color: entry.color } : {}),
        };
      });
      return { ...common, kind: "chart", series };
    }

    throw new Error(`The dashboard API returned an unsupported section in "${section.title}".`);
  });

  let user: DashboardPayload["user"];
  if (value.user !== undefined) {
    if (!isRecord(value.user) || (value.user.name !== undefined && typeof value.user.name !== "string") || (value.user.roleLabel !== undefined && typeof value.user.roleLabel !== "string")) {
      throw new Error("The dashboard API returned an invalid user profile.");
    }
    user = {
      ...(typeof value.user.name === "string" ? { name: value.user.name } : {}),
      ...(typeof value.user.roleLabel === "string" ? { roleLabel: value.user.roleLabel } : {}),
    };
  }

  return { role, metrics, sections, ...(user ? { user } : {}) };
}

function WidgetContent({ widget }: { widget: DashboardWidget }) {
  if (widget.kind === "table") {
    if (!widget.rows.length) return <div className="dashboard-widget-empty">Nothing to show yet.</div>;
    return (
      <div className="dashboard-table-scroll">
        <table className="dashboard-table">
          <thead><tr>{widget.columns.map((column) => <th key={column} scope="col">{column}</th>)}</tr></thead>
          <tbody>
            {widget.rows.map((row, index) => (
              <tr key={`${widget.id}-${index}`}>
                {widget.columns.map((column) => {
                  const value = row[column] ?? row[column.toLowerCase().replaceAll(" ", "_")] ?? "—";
                  return <td key={column}>{String(value)}</td>;
                })}
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    );
  }

  if (widget.kind === "progress") {
    if (!widget.items.length) return <div className="dashboard-widget-empty">No progress data yet.</div>;
    return (
      <div className="dashboard-progress-list">
        {widget.items.map((item) => {
          const max = item.max && item.max > 0 ? item.max : 100;
          const percent = Math.max(0, Math.min(100, (item.value / max) * 100));
          return (
            <div className="dashboard-progress-row" key={item.label}>
              <div><span>{item.label}</span><strong>{item.valueLabel ?? `${Math.round(percent)}%`}</strong></div>
              <div className="dashboard-progress-track" role="progressbar" aria-label={item.label} aria-valuenow={item.value} aria-valuemin={0} aria-valuemax={max}>
                <span style={{ width: `${percent}%` }} />
              </div>
            </div>
          );
        })}
      </div>
    );
  }

  if (widget.kind === "summary") {
    if (!widget.items.length) return <div className="dashboard-widget-empty">No summary data yet.</div>;
    return (
      <div className="dashboard-summary-grid">
        {widget.items.map((item) => (
          <DesignCard as="div" className="dashboard-summary-item" key={item.label}>
            <span>{item.label}</span>
            <strong>{item.value}</strong>
            {item.description && <small>{item.description}</small>}
          </DesignCard>
        ))}
      </div>
    );
  }

  if (widget.kind === "chart") {
    const points = widget.series.flatMap((series) => series.values);
    if (!points.length) return <div className="dashboard-widget-empty">No chart data yet.</div>;
    const max = Math.max(1, ...points.map((point) => point.value));
    return (
      <div className="dashboard-chart">
        <div className="dashboard-chart-bars">
          {widget.series.map((series) => (
            <div className="dashboard-chart-series" key={series.name}>
              {series.values.map((point, index) => (
                <div className="dashboard-chart-point" key={`${point.label}-${index}`}>
                  <div className="dashboard-chart-bar-track">
                    <span style={{ height: `${Math.max(3, (point.value / max) * 100)}%`, ...(series.color ? { backgroundColor: series.color } : {}) }} />
                  </div>
                  <small>{point.label}</small>
                </div>
              ))}
            </div>
          ))}
        </div>
        <div className="dashboard-chart-legend">
          {widget.series.map((series) => <span key={series.name}><i style={series.color ? { backgroundColor: series.color } : undefined} />{series.name}</span>)}
        </div>
      </div>
    );
  }

  if (!widget.items.length) return <div className="dashboard-widget-empty">Nothing here yet.</div>;
  return (
    <ul className="dashboard-activity-list">
      {widget.items.map((item, index) => (
        <li key={`${widget.id}-${index}`}>
          <span className="dashboard-activity-mark"><Activity size={15} /></span>
          <span className="dashboard-activity-copy"><strong>{item.title}</strong>{item.subtitle && <small>{item.subtitle}</small>}</span>
          {item.status && <span className={`dashboard-status${/pending|due|critical|low/i.test(item.status) ? " is-attention" : ""}`}>{item.status}</span>}
          {item.value !== undefined && <strong className="dashboard-item-value">{item.value}</strong>}
        </li>
      ))}
    </ul>
  );
}

function DashboardWidgetCard({ widget }: { widget: DashboardWidget }) {
  return (
    <DesignCard as="section" className={`dashboard-widget dashboard-size-${widget.size ?? "half"}`} id={`widget-${widget.id}`}>
      <div className="dashboard-widget-heading">
        <h2>{widget.title}</h2>
        <a href={`#widget-${widget.id}`} aria-label={`Link to ${widget.title}`}><ArrowRight size={15} /></a>
      </div>
      <WidgetContent widget={widget} />
    </DesignCard>
  );
}

export function DashboardView({ role }: { role: DashboardRole }) {
  const [payload, setPayload] = useState<DashboardPayload | null>(null);
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(true);
  const [search, setSearch] = useState("");
  const [menuOpen, setMenuOpen] = useState(false);
  const config = roleDetails[role];

  const loadDashboard = useCallback(async (signal?: AbortSignal) => {
    setLoading(true);
    setError("");
    setPayload(null);
    const apiBase = process.env.NEXT_PUBLIC_EDU247_API_BASE_URL?.replace(/\/+$/, "") ?? "";

    try {
      const response = await fetch(`${apiBase}/api/v1/dashboard/${role}`, {
        method: "GET",
        headers: { Accept: "application/json" },
        credentials: "include",
        cache: "no-store",
        signal,
      });
      if (!response.ok) {
        throw new Error(`Dashboard API request failed (${response.status}).`);
      }
      const body: unknown = await response.json();
      setPayload(parseDashboardPayload(body, role));
    } catch (cause) {
      if (cause instanceof DOMException && cause.name === "AbortError") return;
      setError(cause instanceof Error ? cause.message : "The dashboard could not be loaded.");
    } finally {
      if (!signal?.aborted) setLoading(false);
    }
  }, [role]);

  useEffect(() => {
    const controller = new AbortController();
    void loadDashboard(controller.signal);
    return () => controller.abort();
  }, [loadDashboard]);

  const normalizedSearch = search.trim().toLowerCase();
  const filteredSections = payload?.sections.filter((section) =>
    section.title.toLowerCase().includes(normalizedSearch),
  ) ?? [];

  return (
    <main className="dashboard-app">
      <aside className={`dashboard-sidebar${menuOpen ? " dashboard-sidebar-open" : ""}`}>
        <Link className="dashboard-brand" href="/" aria-label="EDU247 home">
          <span className="dashboard-brand-mark">
            <Image src={`${basePath}/logoimg/edu247-logo.png`} alt="" width={40} height={40} priority />
          </span>
          <span><strong>EDU247</strong><small>{payload?.user?.roleLabel ?? config.roleLabel}</small></span>
        </Link>
        <nav className="dashboard-side-nav" aria-label={`${config.roleLabel} dashboard`}>
          {config.nav.map(({ label, icon: Icon }, index) => (
            <a
              aria-current={index === 0 ? "page" : undefined}
              className={index === 0 ? "dashboard-nav-active" : ""}
              href={index === 0 ? "#dashboard-home" : `#widget-${label.toLowerCase().replaceAll(/[^a-z0-9]+/g, "-")}`}
              key={label}
              onClick={() => setMenuOpen(false)}
            >
              <Icon size={17} /><span>{label}</span>
            </a>
          ))}
        </nav>
        <div className="dashboard-sidebar-footer"><Shield size={15} /> Your learning space</div>
      </aside>

      <div className="dashboard-main">
        <header className="dashboard-topbar">
          <button
            className="dashboard-menu-button"
            type="button"
            aria-label={menuOpen ? "Close dashboard menu" : "Open dashboard menu"}
            aria-expanded={menuOpen}
            onClick={() => setMenuOpen((open) => !open)}
          >
            {menuOpen ? <X size={19} /> : <Menu size={19} />}
          </button>
          <span className="dashboard-page-title">{config.title}</span>
          <label className="dashboard-search">
            <Search size={16} />
            <input
              aria-label="Search dashboard sections"
              onChange={(event) => setSearch(event.target.value)}
              placeholder="Search dashboard..."
              value={search}
            />
          </label>
          <Link className="dashboard-user" href="/">
            <span className="dashboard-avatar">{(payload?.user?.name ?? config.roleLabel).slice(0, 1).toUpperCase()}</span>
            <span>{payload?.user?.name ?? config.roleLabel}</span>
            <ChevronDown size={15} />
          </Link>
        </header>

        <div className="dashboard-content">
          <section className="dashboard-welcome" id="dashboard-home">
            <div>
              <span className="dashboard-eyebrow">YOUR EDU247 WORKSPACE</span>
              <h1>Welcome back{payload?.user?.name ? `, ${payload.user.name}` : ""}.</h1>
              <p>Your dashboard is powered by your latest platform data.</p>
            </div>
            <button className="dashboard-refresh" type="button" onClick={() => void loadDashboard()} disabled={loading}>
              <RefreshCw size={15} className={loading ? "dashboard-spin" : ""} /> Refresh
            </button>
          </section>

          {loading ? (
            <DesignCard as="section" className="dashboard-loading" role="status" aria-live="polite">
              <LoaderCircle className="dashboard-spin" size={24} />
              <span>Loading your dashboard...</span>
              <div className="dashboard-skeleton-grid" aria-hidden="true">
                {Array.from({ length: 4 }, (_, index) => <i key={index} />)}
              </div>
            </DesignCard>
          ) : error ? (
            <DesignCard as="section" className="dashboard-error" role="alert">
              <span className="dashboard-error-icon"><AlertCircle size={22} /></span>
              <div>
                <h2>Dashboard data is unavailable</h2>
                <p>{error} Expected endpoint: <code>/api/v1/dashboard/{role}</code></p>
                <p>Check that the EDU247 API is running and that this account has access to the dashboard.</p>
                <button type="button" className="dashboard-retry" onClick={() => void loadDashboard()}>
                  <RefreshCw size={15} /> Try again
                </button>
              </div>
            </DesignCard>
          ) : payload ? (
            <>
              {payload.metrics.length > 0 && (
                <section className="dashboard-metrics" aria-label={`${config.roleLabel} metrics`}>
                  {payload.metrics.map((metric) => {
                    const TrendIcon = metric.trend === "down" ? ArrowDownRight : ArrowUpRight;
                    return (
                      <DesignCard className="dashboard-metric" key={metric.id}>
                        <span className="dashboard-metric-icon"><ChartNoAxesCombined size={18} /></span>
                        <div><span>{metric.label}</span><strong>{metric.value}</strong>
                          {metric.change && <small className={`dashboard-trend trend-${metric.trend ?? "neutral"}`}>{metric.trend && <TrendIcon size={12} />}{metric.change}</small>}
                        </div>
                      </DesignCard>
                    );
                  })}
                </section>
              )}

              {filteredSections.length > 0 ? (
                <section className="dashboard-widget-grid" aria-label="Dashboard sections">
                  {filteredSections.map((widget) => <DashboardWidgetCard key={widget.id} widget={widget} />)}
                </section>
              ) : normalizedSearch ? (
                <section className="dashboard-search-empty"><p>No sections match “{search}”.</p></section>
              ) : (
                <DesignCard as="section" className="dashboard-empty">
                  <span><LayoutDashboard size={25} /></span>
                  <h2>{payload.metrics.length ? "No dashboard sections yet" : "Your dashboard is ready for data"}</h2>
                  <p>{payload.metrics.length ? "Metrics are available, but the API returned no dashboard sections yet." : "The API returned no dashboard metrics or sections for this account yet."}</p>
                  <button type="button" className="dashboard-retry" onClick={() => void loadDashboard()}><RefreshCw size={15} /> Refresh data</button>
                </DesignCard>
              )}
            </>
          ) : null}

        </div>
      </div>
    </main>
  );
}
