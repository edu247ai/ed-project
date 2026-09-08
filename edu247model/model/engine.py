from edu247model.model.core.supervisor import Supervisor
from edu247model.skills.loader import load_skill

def build_engine():
    sup = Supervisor()

    sup.register_skill(load_skill("edu247model.skills.anomaly_skill", "AnomalySkill", sup))
    sup.register_skill(load_skill("edu247model.skills.forecast_skill", "ForecastSkill", sup))
    sup.register_skill(load_skill("edu247model.skills.autoscale_skill", "AutoscaleSkill", sup))
    sup.register_skill(load_skill("edu247model.skills.resource_opt_skill", "ResourceOptSkill", sup))
    sup.register_skill(load_skill("edu247model.skills.rag_skill", "RAGSkill", sup))
    sup.register_skill(load_skill("edu247model.skills.rl_skill", "RLSkill", sup))

    return sup
