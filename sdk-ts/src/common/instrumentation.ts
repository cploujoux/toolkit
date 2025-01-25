import { OTLPLogExporter } from "@opentelemetry/exporter-logs-otlp-grpc";
import { OTLPMetricExporter } from "@opentelemetry/exporter-metrics-otlp-grpc";
import { OTLPTraceExporter } from "@opentelemetry/exporter-trace-otlp-grpc";
import { FastifyInstrumentation } from "@opentelemetry/instrumentation-fastify";
import { HttpInstrumentation } from "@opentelemetry/instrumentation-http";
import { Resource } from "@opentelemetry/resources";
import {
  BatchLogRecordProcessor,
  LoggerProvider,
} from "@opentelemetry/sdk-logs";
import {
  MeterProvider,
  PeriodicExportingMetricReader,
} from "@opentelemetry/sdk-metrics";
import { BatchSpanProcessor } from "@opentelemetry/sdk-trace-base";
import { NodeTracerProvider } from "@opentelemetry/sdk-trace-node";
import { SemanticResourceAttributes } from "@opentelemetry/semantic-conventions";
import { FastifyInstance } from "fastify";
import { getAuthenticationHeaders } from "../authentication/authentication.js";
import { getSettings } from "./settings.js";

let tracerProvider: NodeTracerProvider | null = null;
let meterProvider: MeterProvider | null = null;
let loggerProvider: LoggerProvider | null = null;

/**
 * Retrieve authentication headers.
 */
async function authHeaders(): Promise<Record<string, string>> {
  const settings = getSettings();
  const headers = await getAuthenticationHeaders(settings);
  return {
    "x-beamlit-authorization": headers?.["X-Beamlit-Authorization"] || "",
    "x-beamlit-workspace": headers?.["X-Beamlit-Workspace"] || "",
  };
}

/**
 * Initialize and return the LoggerProvider.
 */
export function getLoggerProviderInstance(): LoggerProvider {
  if (!loggerProvider) {
    throw new Error("LoggerProvider is not initialized");
  }
  return loggerProvider;
}

/**
 * Get resource attributes for OpenTelemetry.
 */
function getResourceAttributes(): Record<string, any> {
  const settings = getSettings();
  return {
    [SemanticResourceAttributes.SERVICE_NAME]: settings.name,
    workspace: settings.workspace,
    "service.namespace": settings.workspace,
    "service.workspace": settings.workspace,
  };
}

/**
 * Initialize and return the OTLP Metric Exporter.
 */
async function getMetricExporter(): Promise<OTLPMetricExporter | null> {
  const settings = getSettings();
  if (!settings.enableOpentelemetry) {
    return null;
  }
  return new OTLPMetricExporter({
    headers: await authHeaders(),
  });
}

/**
 * Initialize and return the OTLP Trace Exporter.
 */
async function getTraceExporter(): Promise<OTLPTraceExporter | null> {
  const settings = getSettings();
  if (!settings.enableOpentelemetry) {
    return null;
  }
  return new OTLPTraceExporter({
    headers: await authHeaders(),
  });
}

/**
 * Initialize and return the OTLP Log Exporter.
 */
async function getLogExporter(): Promise<OTLPLogExporter | null> {
  const settings = getSettings();
  if (!settings.enableOpentelemetry) {
    return null;
  }
  return new OTLPLogExporter({
    headers: await authHeaders(),
  });
}

/**
 * Instrument the Fastify application with OpenTelemetry.
 * @param app Fastify instance
 */
export async function instrumentApp(app: FastifyInstance) {
  const settings = getSettings();
  if (!settings.enableOpentelemetry) {
    return;
  }

  const resource = new Resource(getResourceAttributes());

  // Initialize Tracer Provider
  tracerProvider = new NodeTracerProvider({
    resource,
  });
  const traceExporter = await getTraceExporter();
  if (traceExporter) {
    tracerProvider.addSpanProcessor(new BatchSpanProcessor(traceExporter));
  }
  tracerProvider.register();

  // Initialize Meter Provider
  meterProvider = new MeterProvider({
    resource,
    // Add more configurations if needed
  });
  const metricExporter = await getMetricExporter();
  if (metricExporter) {
    meterProvider.addMetricReader(
      new PeriodicExportingMetricReader({
        exporter: metricExporter,
        exportIntervalMillis: 60000,
      })
    );
  }

  // Initialize Logger Provider
  loggerProvider = new LoggerProvider({
    resource,
  });
  const logExporter = await getLogExporter();
  if (logExporter) {
    loggerProvider.addLogRecordProcessor(
      new BatchLogRecordProcessor(logExporter)
    );
  }
  // Instrument Fastify and HTTP
  const fastifyInstrumentation = new FastifyInstrumentation();
  fastifyInstrumentation.enable();

  const httpInstrumentation = new HttpInstrumentation();
  httpInstrumentation.enable();
}

/**
 * Shutdown OpenTelemetry instrumentation.
 */
export function shutdownInstrumentation() {
  if (tracerProvider) {
    tracerProvider
      .shutdown()
      .then(() => {
        console.log("Tracer provider shut down successfully");
      })
      .catch((error: Error) => {
        console.error("Error shutting down tracer provider", error);
      });
  }

  if (meterProvider) {
    meterProvider
      .shutdown()
      .then(() => {
        console.log("Meter provider shut down successfully");
      })
      .catch((error: Error) => {
        console.error("Error shutting down meter provider", error);
      });
  }

  if (loggerProvider) {
    loggerProvider
      .shutdown()
      .then(() => {
        console.log("Logger provider shut down successfully");
      })
      .catch((error: Error) => {
        console.error("Error shutting down logger provider", error);
      });
  }
}
