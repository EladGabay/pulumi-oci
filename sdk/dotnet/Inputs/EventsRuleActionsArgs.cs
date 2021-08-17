// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Oci.Inputs
{

    public sealed class EventsRuleActionsArgs : Pulumi.ResourceArgs
    {
        [Input("actions", required: true)]
        private InputList<Inputs.EventsRuleActionsActionArgs>? _actions;

        /// <summary>
        /// (Updatable) A list of one or more ActionDetails objects.
        /// </summary>
        public InputList<Inputs.EventsRuleActionsActionArgs> Actions
        {
            get => _actions ?? (_actions = new InputList<Inputs.EventsRuleActionsActionArgs>());
            set => _actions = value;
        }

        public EventsRuleActionsArgs()
        {
        }
    }
}